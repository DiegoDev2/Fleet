package main

// FormatUdf - Bash script to format a block device to UDF
// Homepage: https://github.com/JElchison/format-udf

import (
	"fmt"
	
	"os/exec"
)

func installFormatUdf() {
	// Método 1: Descargar y extraer .tar.gz
	formatudf_tar_url := "https://github.com/JElchison/format-udf/archive/refs/tags/1.8.0.tar.gz"
	formatudf_cmd_tar := exec.Command("curl", "-L", formatudf_tar_url, "-o", "package.tar.gz")
	err := formatudf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	formatudf_zip_url := "https://github.com/JElchison/format-udf/archive/refs/tags/1.8.0.zip"
	formatudf_cmd_zip := exec.Command("curl", "-L", formatudf_zip_url, "-o", "package.zip")
	err = formatudf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	formatudf_bin_url := "https://github.com/JElchison/format-udf/archive/refs/tags/1.8.0.bin"
	formatudf_cmd_bin := exec.Command("curl", "-L", formatudf_bin_url, "-o", "binary.bin")
	err = formatudf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	formatudf_src_url := "https://github.com/JElchison/format-udf/archive/refs/tags/1.8.0.src.tar.gz"
	formatudf_cmd_src := exec.Command("curl", "-L", formatudf_src_url, "-o", "source.tar.gz")
	err = formatudf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	formatudf_cmd_direct := exec.Command("./binary")
	err = formatudf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
