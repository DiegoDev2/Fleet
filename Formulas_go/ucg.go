package main

// Ucg - Tool for searching large bodies of source code (like grep)
// Homepage: https://github.com/gvansickle/ucg

import (
	"fmt"
	
	"os/exec"
)

func installUcg() {
	// Método 1: Descargar y extraer .tar.gz
	ucg_tar_url := "https://github.com/gvansickle/ucg/releases/download/0.3.3/universalcodegrep-0.3.3.tar.gz"
	ucg_cmd_tar := exec.Command("curl", "-L", ucg_tar_url, "-o", "package.tar.gz")
	err := ucg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ucg_zip_url := "https://github.com/gvansickle/ucg/releases/download/0.3.3/universalcodegrep-0.3.3.zip"
	ucg_cmd_zip := exec.Command("curl", "-L", ucg_zip_url, "-o", "package.zip")
	err = ucg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ucg_bin_url := "https://github.com/gvansickle/ucg/releases/download/0.3.3/universalcodegrep-0.3.3.bin"
	ucg_cmd_bin := exec.Command("curl", "-L", ucg_bin_url, "-o", "binary.bin")
	err = ucg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ucg_src_url := "https://github.com/gvansickle/ucg/releases/download/0.3.3/universalcodegrep-0.3.3.src.tar.gz"
	ucg_cmd_src := exec.Command("curl", "-L", ucg_src_url, "-o", "source.tar.gz")
	err = ucg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ucg_cmd_direct := exec.Command("./binary")
	err = ucg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: argp-standalone")
exec.Command("latte", "install", "argp-standalone")
}
