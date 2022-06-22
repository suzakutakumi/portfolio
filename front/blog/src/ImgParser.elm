module ImgParser exposing (..)

import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)
import Parser exposing (..)
import Models exposing (..)
type alias Image = 
    { befText : String
    , alt : String
    , url : String
    , aftText : String
    }

parse:String -> Model
parse txt=
    Result.withDefault (div [] [text txt]) <|(Result.map toImg <|run imgParser txt)

toImg:Image->Model
toImg picture=
    div[] [
        text picture.befText,
        img [src picture.url, alt picture.alt] [],
        parse picture.aftText
    ]

imgParser:Parser Image
imgParser =
    succeed Image
        |= (getChompedString <| chompUntil "![")
        |. symbol "["
        |= (getChompedString <| chompUntil "](")
        |. symbol "]("
        |= (getChompedString <| chompUntil ")")
        |. symbol ")"
        |= (getChompedString <| chompUntilEndOr "\n")