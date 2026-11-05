package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"threat-api/internal/model"
	"threat-api/internal/validator"

	"github.com/jackc/pgx/v5/pgconn"
)

func CreateThreat(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var threat model.Threat

		err := json.NewDecoder(r.Body).Decode(&threat)
		if err != nil {
			WriteError(w, http.StatusBadRequest,
				"Invalid JSON", "Failed to decode request body")
			return
		}

		if threat.Severity == "" {
			threat.Severity = "low"
		}

		if !validator.ValidateThreat(threat) {
			WriteError(w, http.StatusBadRequest,
				"Invalid threat data",
				"MD5, name or severity is invalid")
			return
		}

		threat.DetectedTime = time.Now()

		_, err = db.ExecContext(
			r.Context(),
			`INSERT INTO threats
			 (md5, name, description, severity, detected_time)
			 VALUES ($1, $2, $3, $4, $5)`,
			threat.MD5,
			threat.Name,
			threat.Description,
			threat.Severity,
			threat.DetectedTime,
		)

		if err != nil {
			var pgErr *pgconn.PgError

			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				WriteError(w, http.StatusConflict,
					"Threat already exists",
					"A threat with this MD5 already exists")
				return
			}

			fmt.Println("Database error:", err)
			WriteError(w, http.StatusInternalServerError,
				"Database error", "Failed to create threat")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(threat); err != nil {
			fmt.Println("Error encoding JSON:", err)
		}
	}
}
