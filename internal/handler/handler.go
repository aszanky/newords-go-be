package handler

import (
	"net/http"

	"github.com/aszanky/newords-go-be/internal/models"
	"github.com/aszanky/newords-go-be/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	usecase  usecase.Usecase
	router   *gin.Engine
	validate *validator.Validate
}

func NewHandler(
	uc usecase.Usecase,
) *Handler {
	r := gin.Default()
	val := validator.New()

	return &Handler{
		usecase:  uc,
		router:   r,
		validate: val,
	}
}

func (h *Handler) AddNewWord(c *gin.Context) {
	var input models.Word

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.AddNewWords(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"word": input,
	})
}

func (r *Handler) Register() {
	r.router.POST("/words/add", r.AddNewWord)
}

// Start HTTP Server
func (r *Handler) Start(port string) error {
	// url := host + port

	//solve issue You trusted all proxies, this is NOT safe. We recommend you to set a value.

	r.router.ForwardedByClientIP = true
	r.router.SetTrustedProxies([]string{"0.0.0.0:6010"})
	r.Register()

	return r.router.Run(port)
}
