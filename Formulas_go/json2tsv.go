package main

// Json2tsv - JSON to TSV converter
// Homepage: https://codemadness.org/json2tsv.html

import (
	"fmt"
	
	"os/exec"
)

func installJson2tsv() {
	// Método 1: Descargar y extraer .tar.gz
	json2tsv_tar_url := "https://codemadness.org/releases/json2tsv/json2tsv-1.2.tar.gz"
	json2tsv_cmd_tar := exec.Command("curl", "-L", json2tsv_tar_url, "-o", "package.tar.gz")
	err := json2tsv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	json2tsv_zip_url := "https://codemadness.org/releases/json2tsv/json2tsv-1.2.zip"
	json2tsv_cmd_zip := exec.Command("curl", "-L", json2tsv_zip_url, "-o", "package.zip")
	err = json2tsv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	json2tsv_bin_url := "https://codemadness.org/releases/json2tsv/json2tsv-1.2.bin"
	json2tsv_cmd_bin := exec.Command("curl", "-L", json2tsv_bin_url, "-o", "binary.bin")
	err = json2tsv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	json2tsv_src_url := "https://codemadness.org/releases/json2tsv/json2tsv-1.2.src.tar.gz"
	json2tsv_cmd_src := exec.Command("curl", "-L", json2tsv_src_url, "-o", "source.tar.gz")
	err = json2tsv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	json2tsv_cmd_direct := exec.Command("./binary")
	err = json2tsv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
