#!/bin/bash

LICENSE="// Copyright (C) 2024 Fleet Inc.
//
// Licensed under the Apache License, Version 2.0 (the \"License\");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an \"AS IS\" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formulas
"

# Recorrer carpetas específicas: CLI, Formulas, handlers
for file in $(find ./cli ./formulas ./handlers ./lib -name "*.go"); do
    if ! grep -q "Copyright (C) 2024 Fleet Inc." "$file"; then
        echo -e "$LICENSE\n\n$(cat $file)" > $file
        echo "Licencia añadida a $file"
    else
        echo "Licencia ya existe en $file"
    fi
done
