package main

// Libxtst - X.Org: Client API for the XTEST & RECORD extensions
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxtst() {
	// Método 1: Descargar y extraer .tar.gz
	libxtst_tar_url := "https://www.x.org/archive/individual/lib/libXtst-1.2.5.tar.gz"
	libxtst_cmd_tar := exec.Command("curl", "-L", libxtst_tar_url, "-o", "package.tar.gz")
	err := libxtst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxtst_zip_url := "https://www.x.org/archive/individual/lib/libXtst-1.2.5.zip"
	libxtst_cmd_zip := exec.Command("curl", "-L", libxtst_zip_url, "-o", "package.zip")
	err = libxtst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxtst_bin_url := "https://www.x.org/archive/individual/lib/libXtst-1.2.5.bin"
	libxtst_cmd_bin := exec.Command("curl", "-L", libxtst_bin_url, "-o", "binary.bin")
	err = libxtst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxtst_src_url := "https://www.x.org/archive/individual/lib/libXtst-1.2.5.src.tar.gz"
	libxtst_cmd_src := exec.Command("curl", "-L", libxtst_src_url, "-o", "source.tar.gz")
	err = libxtst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxtst_cmd_direct := exec.Command("./binary")
	err = libxtst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxi")
	exec.Command("latte", "install", "libxi").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
}
