package main

// Libxkbfile - X.Org: XKB file handling routines
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxkbfile() {
	// Método 1: Descargar y extraer .tar.gz
	libxkbfile_tar_url := "https://www.x.org/archive/individual/lib/libxkbfile-1.1.3.tar.xz"
	libxkbfile_cmd_tar := exec.Command("curl", "-L", libxkbfile_tar_url, "-o", "package.tar.gz")
	err := libxkbfile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxkbfile_zip_url := "https://www.x.org/archive/individual/lib/libxkbfile-1.1.3.tar.xz"
	libxkbfile_cmd_zip := exec.Command("curl", "-L", libxkbfile_zip_url, "-o", "package.zip")
	err = libxkbfile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxkbfile_bin_url := "https://www.x.org/archive/individual/lib/libxkbfile-1.1.3.tar.xz"
	libxkbfile_cmd_bin := exec.Command("curl", "-L", libxkbfile_bin_url, "-o", "binary.bin")
	err = libxkbfile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxkbfile_src_url := "https://www.x.org/archive/individual/lib/libxkbfile-1.1.3.tar.xz"
	libxkbfile_cmd_src := exec.Command("curl", "-L", libxkbfile_src_url, "-o", "source.tar.gz")
	err = libxkbfile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxkbfile_cmd_direct := exec.Command("./binary")
	err = libxkbfile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}
