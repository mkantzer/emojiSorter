package api

// func (a *Server) testLog(c *gin.Context) {
// 	a.Deps.Logger.Info("hey look at me!")
// 	c.String(http.StatusOK, "Hewdy %s", "bruh")
// }

// func (a *Server) getEmojiByName(c *gin.Context) {
// 	name := c.Param("name")
// 	emoji, err := a.Deps.Database.GetEmojiByName(c.Request.Context(), name)
// 	if err != nil {
// 		switch err {
// 		case apperrors.ErrEmojiNotFound:
// 			c.String(http.StatusNotFound, err.Error())
// 		default:
// 			c.String(http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, emoji)

// }

// func (a *Server) getAllEmoji(c *gin.Context) {
// 	emoji, err := a.Deps.Database.GetAllEmoji(c.Request.Context())
// 	if err != nil {
// 		switch err {
// 		case apperrors.ErrEmojiNotFound:
// 			c.String(http.StatusNotFound, err.Error())
// 		default:
// 			c.String(http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, emoji)
// }
