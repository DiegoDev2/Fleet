package main

// BdwGc - Garbage collector for C and C++
// Homepage: https://www.hboehm.info/gc/

import (
	"fmt"
	
	"os/exec"
)

func installBdwGc() {
	// Método 1: Descargar y extraer .tar.gz
	bdwgc_tar_url := "https://github.com/ivmai/bdwgc/releases/download/v8.2.6/gc-8.2.6.tar.gz"
	bdwgc_cmd_tar := exec.Command("curl", "-L", bdwgc_tar_url, "-o", "package.tar.gz")
	err := bdwgc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bdwgc_zip_url := "https://github.com/ivmai/bdwgc/releases/download/v8.2.6/gc-8.2.6.zip"
	bdwgc_cmd_zip := exec.Command("curl", "-L", bdwgc_zip_url, "-o", "package.zip")
	err = bdwgc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bdwgc_bin_url := "https://github.com/ivmai/bdwgc/releases/download/v8.2.6/gc-8.2.6.bin"
	bdwgc_cmd_bin := exec.Command("curl", "-L", bdwgc_bin_url, "-o", "binary.bin")
	err = bdwgc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bdwgc_src_url := "https://github.com/ivmai/bdwgc/releases/download/v8.2.6/gc-8.2.6.src.tar.gz"
	bdwgc_cmd_src := exec.Command("curl", "-L", bdwgc_src_url, "-o", "source.tar.gz")
	err = bdwgc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bdwgc_cmd_direct := exec.Command("./binary")
	err = bdwgc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libatomic_ops")
	exec.Command("latte", "install", "libatomic_ops").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
