package main

// Plantuml - Draw UML diagrams
// Homepage: https://plantuml.com/

import (
	"fmt"
	
	"os/exec"
)

func installPlantuml() {
	// Método 1: Descargar y extraer .tar.gz
	plantuml_tar_url := "https://github.com/plantuml/plantuml/releases/download/v1.2024.6/plantuml-1.2024.6.jar"
	plantuml_cmd_tar := exec.Command("curl", "-L", plantuml_tar_url, "-o", "package.tar.gz")
	err := plantuml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plantuml_zip_url := "https://github.com/plantuml/plantuml/releases/download/v1.2024.6/plantuml-1.2024.6.jar"
	plantuml_cmd_zip := exec.Command("curl", "-L", plantuml_zip_url, "-o", "package.zip")
	err = plantuml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plantuml_bin_url := "https://github.com/plantuml/plantuml/releases/download/v1.2024.6/plantuml-1.2024.6.jar"
	plantuml_cmd_bin := exec.Command("curl", "-L", plantuml_bin_url, "-o", "binary.bin")
	err = plantuml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plantuml_src_url := "https://github.com/plantuml/plantuml/releases/download/v1.2024.6/plantuml-1.2024.6.jar"
	plantuml_cmd_src := exec.Command("curl", "-L", plantuml_src_url, "-o", "source.tar.gz")
	err = plantuml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plantuml_cmd_direct := exec.Command("./binary")
	err = plantuml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
