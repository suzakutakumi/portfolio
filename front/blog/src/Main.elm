module Main exposing (..)

-- Press buttons to increment and decrement a counter.
--
-- Read how it works:
--   https://guide.elm-lang.org/architecture/buttons.html
--

import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)
import Regex
import ImgParser
import Models exposing(..)


-- MAIN


main =
    Browser.sandbox { init = init, update = update, view = view }


-- UPDATE





update : Msg -> Model -> Model
update msg model =
    case msg of
        Change rows ->
            div [] <|
                List.map header <|
                    List.map  (String.lines rows)


header : String -> Html Msg
header row =
    let
        cnt =
            shapeCount row
    in
    if cnt > 0 && Regex.contains (regStr "^#* ") row then
        chooseHeader cnt [] [ text <| String.dropLeft cnt row ]

    else
        div [] [ text row ]


regStr : String->Regex.Regex
regStr reg =
    Maybe.withDefault Regex.never (Regex.fromString reg)


shapeCount : String -> Int
shapeCount row =
    if String.startsWith "#" row == True then
        (shapeCount <| String.dropLeft 1 row) + 1

    else
        0


chooseHeader : Int -> (List (Attribute Msg) -> List (Html Msg) -> Html Msg)
chooseHeader num =
    case num of
        0 ->
            div

        1 ->
            h1

        2 ->
            h2

        3 ->
            h3

        4 ->
            h4

        5 ->
            h5

        _ ->
            h6

-- VIEW


view : Model -> Html Msg
view model =
    div []
        [ textarea [ onInput Change ] []
        , div [] [ model ]
        ]
