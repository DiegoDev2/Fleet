package main

// Expect - Program that can automate interactive applications
// Homepage: https://core.tcl-lang.org/expect/index

import (
	"fmt"
	
	"os/exec"
)

func installExpect() {
	// Método 1: Descargar y extraer .tar.gz
	expect_tar_url := "https://downloads.sourceforge.net/project/expect/Expect/5.45.4/expect5.45.4.tar.gz"
	expect_cmd_tar := exec.Command("curl", "-L", expect_tar_url, "-o", "package.tar.gz")
	err := expect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	expect_zip_url := "https://downloads.sourceforge.net/project/expect/Expect/5.45.4/expect5.45.4.zip"
	expect_cmd_zip := exec.Command("curl", "-L", expect_zip_url, "-o", "package.zip")
	err = expect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	expect_bin_url := "https://downloads.sourceforge.net/project/expect/Expect/5.45.4/expect5.45.4.bin"
	expect_cmd_bin := exec.Command("curl", "-L", expect_bin_url, "-o", "binary.bin")
	err = expect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	expect_src_url := "https://downloads.sourceforge.net/project/expect/Expect/5.45.4/expect5.45.4.src.tar.gz"
	expect_cmd_src := exec.Command("curl", "-L", expect_src_url, "-o", "source.tar.gz")
	err = expect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	expect_cmd_direct := exec.Command("./binary")
	err = expect_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
}
