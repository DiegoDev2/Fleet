package main

// Blazegraph - Graph database supporting RDF data model, Sesame, and Blueprint APIs
// Homepage: https://www.blazegraph.com/

import (
	"fmt"
	
	"os/exec"
)

func installBlazegraph() {
	// Método 1: Descargar y extraer .tar.gz
	blazegraph_tar_url := "https://github.com/blazegraph/database/releases/download/BLAZEGRAPH_RELEASE_2_1_5/blazegraph.jar"
	blazegraph_cmd_tar := exec.Command("curl", "-L", blazegraph_tar_url, "-o", "package.tar.gz")
	err := blazegraph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blazegraph_zip_url := "https://github.com/blazegraph/database/releases/download/BLAZEGRAPH_RELEASE_2_1_5/blazegraph.jar"
	blazegraph_cmd_zip := exec.Command("curl", "-L", blazegraph_zip_url, "-o", "package.zip")
	err = blazegraph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blazegraph_bin_url := "https://github.com/blazegraph/database/releases/download/BLAZEGRAPH_RELEASE_2_1_5/blazegraph.jar"
	blazegraph_cmd_bin := exec.Command("curl", "-L", blazegraph_bin_url, "-o", "binary.bin")
	err = blazegraph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blazegraph_src_url := "https://github.com/blazegraph/database/releases/download/BLAZEGRAPH_RELEASE_2_1_5/blazegraph.jar"
	blazegraph_cmd_src := exec.Command("curl", "-L", blazegraph_src_url, "-o", "source.tar.gz")
	err = blazegraph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blazegraph_cmd_direct := exec.Command("./binary")
	err = blazegraph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
