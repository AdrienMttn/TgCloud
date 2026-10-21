package models

type Upload struct
{
	URL string `json:"url" binding:"required"`
}