package main

// Csview - High performance csv viewer for cli
// Homepage: https://github.com/wfxr/csview

import (
	"fmt"
	
	"os/exec"
)

func installCsview() {
	// Método 1: Descargar y extraer .tar.gz
	csview_tar_url := "https://github.com/wfxr/csview/archive/refs/tags/v1.3.3.tar.gz"
	csview_cmd_tar := exec.Command("curl", "-L", csview_tar_url, "-o", "package.tar.gz")
	err := csview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csview_zip_url := "https://github.com/wfxr/csview/archive/refs/tags/v1.3.3.zip"
	csview_cmd_zip := exec.Command("curl", "-L", csview_zip_url, "-o", "package.zip")
	err = csview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csview_bin_url := "https://github.com/wfxr/csview/archive/refs/tags/v1.3.3.bin"
	csview_cmd_bin := exec.Command("curl", "-L", csview_bin_url, "-o", "binary.bin")
	err = csview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csview_src_url := "https://github.com/wfxr/csview/archive/refs/tags/v1.3.3.src.tar.gz"
	csview_cmd_src := exec.Command("curl", "-L", csview_src_url, "-o", "source.tar.gz")
	err = csview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csview_cmd_direct := exec.Command("./binary")
	err = csview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
