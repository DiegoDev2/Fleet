package main

// Apachetop - Top-like display of Apache log
// Homepage: https://github.com/tessus/apachetop

import (
	"fmt"
	
	"os/exec"
)

func installApachetop() {
	// Método 1: Descargar y extraer .tar.gz
	apachetop_tar_url := "https://github.com/tessus/apachetop/releases/download/0.23.2/apachetop-0.23.2.tar.gz"
	apachetop_cmd_tar := exec.Command("curl", "-L", apachetop_tar_url, "-o", "package.tar.gz")
	err := apachetop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachetop_zip_url := "https://github.com/tessus/apachetop/releases/download/0.23.2/apachetop-0.23.2.zip"
	apachetop_cmd_zip := exec.Command("curl", "-L", apachetop_zip_url, "-o", "package.zip")
	err = apachetop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachetop_bin_url := "https://github.com/tessus/apachetop/releases/download/0.23.2/apachetop-0.23.2.bin"
	apachetop_cmd_bin := exec.Command("curl", "-L", apachetop_bin_url, "-o", "binary.bin")
	err = apachetop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachetop_src_url := "https://github.com/tessus/apachetop/releases/download/0.23.2/apachetop-0.23.2.src.tar.gz"
	apachetop_cmd_src := exec.Command("curl", "-L", apachetop_src_url, "-o", "source.tar.gz")
	err = apachetop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachetop_cmd_direct := exec.Command("./binary")
	err = apachetop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: adns")
	exec.Command("latte", "install", "adns").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
