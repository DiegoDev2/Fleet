package main

// Libbs2b - Bauer stereophonic-to-binaural DSP
// Homepage: https://bs2b.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibbs2b() {
	// Método 1: Descargar y extraer .tar.gz
	libbs2b_tar_url := "https://downloads.sourceforge.net/project/bs2b/libbs2b/3.1.0/libbs2b-3.1.0.tar.gz"
	libbs2b_cmd_tar := exec.Command("curl", "-L", libbs2b_tar_url, "-o", "package.tar.gz")
	err := libbs2b_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbs2b_zip_url := "https://downloads.sourceforge.net/project/bs2b/libbs2b/3.1.0/libbs2b-3.1.0.zip"
	libbs2b_cmd_zip := exec.Command("curl", "-L", libbs2b_zip_url, "-o", "package.zip")
	err = libbs2b_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbs2b_bin_url := "https://downloads.sourceforge.net/project/bs2b/libbs2b/3.1.0/libbs2b-3.1.0.bin"
	libbs2b_cmd_bin := exec.Command("curl", "-L", libbs2b_bin_url, "-o", "binary.bin")
	err = libbs2b_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbs2b_src_url := "https://downloads.sourceforge.net/project/bs2b/libbs2b/3.1.0/libbs2b-3.1.0.src.tar.gz"
	libbs2b_cmd_src := exec.Command("curl", "-L", libbs2b_src_url, "-o", "source.tar.gz")
	err = libbs2b_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbs2b_cmd_direct := exec.Command("./binary")
	err = libbs2b_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
