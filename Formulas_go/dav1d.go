package main

// Dav1d - AV1 decoder targeted to be small and fast
// Homepage: https://code.videolan.org/videolan/dav1d

import (
	"fmt"
	
	"os/exec"
)

func installDav1d() {
	// Método 1: Descargar y extraer .tar.gz
	dav1d_tar_url := "https://code.videolan.org/videolan/dav1d/-/archive/1.4.3/dav1d-1.4.3.tar.bz2"
	dav1d_cmd_tar := exec.Command("curl", "-L", dav1d_tar_url, "-o", "package.tar.gz")
	err := dav1d_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dav1d_zip_url := "https://code.videolan.org/videolan/dav1d/-/archive/1.4.3/dav1d-1.4.3.tar.bz2"
	dav1d_cmd_zip := exec.Command("curl", "-L", dav1d_zip_url, "-o", "package.zip")
	err = dav1d_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dav1d_bin_url := "https://code.videolan.org/videolan/dav1d/-/archive/1.4.3/dav1d-1.4.3.tar.bz2"
	dav1d_cmd_bin := exec.Command("curl", "-L", dav1d_bin_url, "-o", "binary.bin")
	err = dav1d_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dav1d_src_url := "https://code.videolan.org/videolan/dav1d/-/archive/1.4.3/dav1d-1.4.3.tar.bz2"
	dav1d_cmd_src := exec.Command("curl", "-L", dav1d_src_url, "-o", "source.tar.gz")
	err = dav1d_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dav1d_cmd_direct := exec.Command("./binary")
	err = dav1d_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}
