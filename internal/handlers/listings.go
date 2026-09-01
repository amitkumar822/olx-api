package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/amitkumar822/olx-api/internal/middleware"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListingHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewListingHandler(db *sql.DB, logger *slog.Logger) *ListingHandler {
	return &ListingHandler{
		db:     db,
		logger: logger,
	}
}

// (lh ListingHandler) this function called method reciver
func (lh ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	// request scoped context
	ctx := r.Context()

	rows, err := lh.db.QueryContext(ctx,
		`SELECT id, title, description, price, city, created_at
			FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)
	if err != nil {
		// log.Printf("query: %v", err)
		lh.logger.Error("listins query error", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	listings := []listing{}

	for rows.Next() {
		var l listing
		if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
			lh.logger.Error("rows scan error", "err", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		lh.logger.Info("listings fetched", "total", len(listings))

		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		lh.logger.Error("rows error", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(listings)

}

func (lh ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// request scoped context
	ctx := r.Context()
	// requestId := ctx.Value("requestCtxId").(string)
	requestId := middleware.RequestIDFromContext(ctx)
	id := r.PathValue("id")

	// lh.logger.Debug("debug log", "listing_id", id)
	// lh.logger.Info("starting query", "listing_id", id)
	// lh.logger.Warn("warn log", "listing_id", id)

	_, err := lh.db.ExecContext(ctx,
		`DELETE FROM listings1 WHERE id =$1`, id)

	if err != nil {
		lh.logger.Error("delete failed", "listing_id", id, "request_id", requestId, "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
