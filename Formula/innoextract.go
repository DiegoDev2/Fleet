package main

// Innoextract - Tool to unpack installers created by Inno Setup
// Homepage: https://constexpr.org/innoextract/

import (
	"fmt"
	
	"os/exec"
)

func installInnoextract() {
	// Método 1: Descargar y extraer .tar.gz
	innoextract_tar_url := "https://constexpr.org/innoextract/files/innoextract-1.9.tar.gz"
	innoextract_cmd_tar := exec.Command("curl", "-L", innoextract_tar_url, "-o", "package.tar.gz")
	err := innoextract_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	innoextract_zip_url := "https://constexpr.org/innoextract/files/innoextract-1.9.zip"
	innoextract_cmd_zip := exec.Command("curl", "-L", innoextract_zip_url, "-o", "package.zip")
	err = innoextract_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	innoextract_bin_url := "https://constexpr.org/innoextract/files/innoextract-1.9.bin"
	innoextract_cmd_bin := exec.Command("curl", "-L", innoextract_bin_url, "-o", "binary.bin")
	err = innoextract_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	innoextract_src_url := "https://constexpr.org/innoextract/files/innoextract-1.9.src.tar.gz"
	innoextract_cmd_src := exec.Command("curl", "-L", innoextract_src_url, "-o", "source.tar.gz")
	err = innoextract_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	innoextract_cmd_direct := exec.Command("./binary")
	err = innoextract_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
