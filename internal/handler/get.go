package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"threat-api/internal/model"
	"threat-api/internal/validator"
)

func GetThreat(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		md5 := r.PathValue("md5")

		if len(md5) != 32 || !validator.CharValidate(md5) {
			WriteError(w, http.StatusBadRequest,
				"Invalid MD5",
				"MD5 must contain exactly 32 hexadecimal characters")
			return
		}

		row := db.QueryRowContext(
			r.Context(),
			`SELECT md5, name, description, severity, detected_time
			 FROM threats
			 WHERE md5 = $1`,
			md5,
		)

		var threat model.Threat

		err := row.Scan(
			&threat.MD5,
			&threat.Name,
			&threat.Description,
			&threat.Severity,
			&threat.DetectedTime,
		)

		if err == sql.ErrNoRows {
			WriteError(w, http.StatusNotFound,
				"Threat not found",
				"No threat with the specified MD5 was found")
			return
		}

		if err != nil {
			fmt.Println("Database error:", err)
			WriteError(w, http.StatusInternalServerError,
				"Database error", "Failed to retrieve threat")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(threat); err != nil {
			fmt.Println("Error encoding JSON:", err)
		}
	}
}
