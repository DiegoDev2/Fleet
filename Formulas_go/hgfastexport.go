package main

// HgFastExport - Fast Mercurial to Git converter
// Homepage: https://repo.or.cz/fast-export.git

import (
	"fmt"
	
	"os/exec"
)

func installHgFastExport() {
	// Método 1: Descargar y extraer .tar.gz
	hgfastexport_tar_url := "https://github.com/frej/fast-export/archive/refs/tags/v231118.tar.gz"
	hgfastexport_cmd_tar := exec.Command("curl", "-L", hgfastexport_tar_url, "-o", "package.tar.gz")
	err := hgfastexport_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hgfastexport_zip_url := "https://github.com/frej/fast-export/archive/refs/tags/v231118.zip"
	hgfastexport_cmd_zip := exec.Command("curl", "-L", hgfastexport_zip_url, "-o", "package.zip")
	err = hgfastexport_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hgfastexport_bin_url := "https://github.com/frej/fast-export/archive/refs/tags/v231118.bin"
	hgfastexport_cmd_bin := exec.Command("curl", "-L", hgfastexport_bin_url, "-o", "binary.bin")
	err = hgfastexport_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hgfastexport_src_url := "https://github.com/frej/fast-export/archive/refs/tags/v231118.src.tar.gz"
	hgfastexport_cmd_src := exec.Command("curl", "-L", hgfastexport_src_url, "-o", "source.tar.gz")
	err = hgfastexport_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hgfastexport_cmd_direct := exec.Command("./binary")
	err = hgfastexport_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mercurial")
exec.Command("latte", "install", "mercurial")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
