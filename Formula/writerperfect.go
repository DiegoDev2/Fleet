package main

// Writerperfect - Library for importing WordPerfect documents
// Homepage: https://sourceforge.net/p/libwpd/wiki/writerperfect/

import (
	"fmt"
	
	"os/exec"
)

func installWriterperfect() {
	// Método 1: Descargar y extraer .tar.gz
	writerperfect_tar_url := "https://downloads.sourceforge.net/project/libwpd/writerperfect/writerperfect-0.9.6/writerperfect-0.9.6.tar.xz"
	writerperfect_cmd_tar := exec.Command("curl", "-L", writerperfect_tar_url, "-o", "package.tar.gz")
	err := writerperfect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	writerperfect_zip_url := "https://downloads.sourceforge.net/project/libwpd/writerperfect/writerperfect-0.9.6/writerperfect-0.9.6.tar.xz"
	writerperfect_cmd_zip := exec.Command("curl", "-L", writerperfect_zip_url, "-o", "package.zip")
	err = writerperfect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	writerperfect_bin_url := "https://downloads.sourceforge.net/project/libwpd/writerperfect/writerperfect-0.9.6/writerperfect-0.9.6.tar.xz"
	writerperfect_cmd_bin := exec.Command("curl", "-L", writerperfect_bin_url, "-o", "binary.bin")
	err = writerperfect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	writerperfect_src_url := "https://downloads.sourceforge.net/project/libwpd/writerperfect/writerperfect-0.9.6/writerperfect-0.9.6.tar.xz"
	writerperfect_cmd_src := exec.Command("curl", "-L", writerperfect_src_url, "-o", "source.tar.gz")
	err = writerperfect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	writerperfect_cmd_direct := exec.Command("./binary")
	err = writerperfect_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libodfgen")
	exec.Command("latte", "install", "libodfgen").Run()
	fmt.Println("Instalando dependencia: libwpd")
	exec.Command("latte", "install", "libwpd").Run()
	fmt.Println("Instalando dependencia: libwpg")
	exec.Command("latte", "install", "libwpg").Run()
	fmt.Println("Instalando dependencia: libwps")
	exec.Command("latte", "install", "libwps").Run()
}
