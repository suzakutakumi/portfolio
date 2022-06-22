module Test exposing (..)

-- Press buttons to increment and decrement a counter.
--
-- Read how it works:
--   https://guide.elm-lang.org/architecture/buttons.html
--

import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)
import Parser exposing (..)
import Regex
import List exposing (foldl)



-- MAIN


main =
    Browser.sandbox { init = init, update = update, view = view }



-- MODEL


type alias Model =
    Html Msg


init : Model
init =
    div [] []



-- UPDATE


type Msg
    = Change String


update : Msg -> Model -> Model
update msg model =
    case msg of
        Change rows ->
            div [] <| List.map image (String.lines rows)

-- VIEW


view : Model -> Html Msg
view model =
    div []
        [ textarea [ onInput Change ] []
        , div [] [ model ]
        ]
