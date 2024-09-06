package main

// Csvtk - Cross-platform, efficient and practical CSV/TSV toolkit in Golang
// Homepage: https://bioinf.shenwei.me/csvtk

import (
	"fmt"
	
	"os/exec"
)

func installCsvtk() {
	// Método 1: Descargar y extraer .tar.gz
	csvtk_tar_url := "https://github.com/shenwei356/csvtk/archive/refs/tags/v0.30.0.tar.gz"
	csvtk_cmd_tar := exec.Command("curl", "-L", csvtk_tar_url, "-o", "package.tar.gz")
	err := csvtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csvtk_zip_url := "https://github.com/shenwei356/csvtk/archive/refs/tags/v0.30.0.zip"
	csvtk_cmd_zip := exec.Command("curl", "-L", csvtk_zip_url, "-o", "package.zip")
	err = csvtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csvtk_bin_url := "https://github.com/shenwei356/csvtk/archive/refs/tags/v0.30.0.bin"
	csvtk_cmd_bin := exec.Command("curl", "-L", csvtk_bin_url, "-o", "binary.bin")
	err = csvtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csvtk_src_url := "https://github.com/shenwei356/csvtk/archive/refs/tags/v0.30.0.src.tar.gz"
	csvtk_cmd_src := exec.Command("curl", "-L", csvtk_src_url, "-o", "source.tar.gz")
	err = csvtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csvtk_cmd_direct := exec.Command("./binary")
	err = csvtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
