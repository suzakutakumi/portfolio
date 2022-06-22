module Models exposing (..)

import Html exposing(..)

-- MODEL
type Msg
    = Change String

type alias Model =
    Html Msg


init : Model
init =
    div [] []


