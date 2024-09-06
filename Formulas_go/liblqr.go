package main

// Liblqr - C/C++ seam carving library
// Homepage: https://liblqr.wikidot.com/

import (
	"fmt"
	
	"os/exec"
)

func installLiblqr() {
	// Método 1: Descargar y extraer .tar.gz
	liblqr_tar_url := "https://github.com/carlobaldassi/liblqr/archive/refs/tags/v0.4.3.tar.gz"
	liblqr_cmd_tar := exec.Command("curl", "-L", liblqr_tar_url, "-o", "package.tar.gz")
	err := liblqr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblqr_zip_url := "https://github.com/carlobaldassi/liblqr/archive/refs/tags/v0.4.3.zip"
	liblqr_cmd_zip := exec.Command("curl", "-L", liblqr_zip_url, "-o", "package.zip")
	err = liblqr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblqr_bin_url := "https://github.com/carlobaldassi/liblqr/archive/refs/tags/v0.4.3.bin"
	liblqr_cmd_bin := exec.Command("curl", "-L", liblqr_bin_url, "-o", "binary.bin")
	err = liblqr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblqr_src_url := "https://github.com/carlobaldassi/liblqr/archive/refs/tags/v0.4.3.src.tar.gz"
	liblqr_cmd_src := exec.Command("curl", "-L", liblqr_src_url, "-o", "source.tar.gz")
	err = liblqr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblqr_cmd_direct := exec.Command("./binary")
	err = liblqr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
