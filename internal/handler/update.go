package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"threat-api/internal/model"
	"threat-api/internal/validator"
)

func UpdateThreat(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		md5 := r.PathValue("md5")

		if len(md5) != 32 || !validator.CharValidate(md5) {
			WriteError(w, http.StatusBadRequest,
				"Invalid MD5",
				"MD5 must contain exactly 32 hexadecimal characters")
			return
		}

		var request struct {
			Name        string  `json:"name"`
			Description *string `json:"description"`
			Severity    *string `json:"severity"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			WriteError(w, http.StatusBadRequest,
				"Invalid JSON", "Failed to decode request body")
			return
		}

		// Сначала получаем существующую запись
		var threat model.Threat

		err := db.QueryRowContext(
			r.Context(),
			`SELECT md5, name, description, severity, detected_time
			 FROM threats WHERE md5 = $1`,
			md5,
		).Scan(
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

		// Обновляем переданные поля
		threat.Name = request.Name

		if request.Description != nil {
			threat.Description = *request.Description
		}

		if request.Severity != nil {
			threat.Severity = *request.Severity
		}

		if !validator.ValidateThreat(threat) {
			WriteError(w, http.StatusBadRequest,
				"Invalid threat data",
				"Name or severity is invalid")
			return
		}

		result, err := db.ExecContext(
			r.Context(),
			`UPDATE threats
			 SET name = $1,
			     description = $2,
			     severity = $3
			 WHERE md5 = $4`,
			threat.Name,
			threat.Description,
			threat.Severity,
			threat.MD5,
		)

		if err != nil {
			fmt.Println("Database error:", err)
			WriteError(w, http.StatusInternalServerError,
				"Database error", "Failed to update threat")
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			WriteError(w, http.StatusInternalServerError,
				"Database error", "Failed to check update result")
			return
		}

		if rowsAffected == 0 {
			WriteError(w, http.StatusNotFound,
				"Threat not found",
				"No threat with the specified MD5 was found")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(threat); err != nil {
			fmt.Println("Error encoding JSON:", err)
		}
	}
}
