package main

// Grep - GNU grep, egrep and fgrep
// Homepage: https://www.gnu.org/software/grep/

import (
	"fmt"
	
	"os/exec"
)

func installGrep() {
	// Método 1: Descargar y extraer .tar.gz
	grep_tar_url := "https://ftp.gnu.org/gnu/grep/grep-3.11.tar.xz"
	grep_cmd_tar := exec.Command("curl", "-L", grep_tar_url, "-o", "package.tar.gz")
	err := grep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grep_zip_url := "https://ftp.gnu.org/gnu/grep/grep-3.11.tar.xz"
	grep_cmd_zip := exec.Command("curl", "-L", grep_zip_url, "-o", "package.zip")
	err = grep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grep_bin_url := "https://ftp.gnu.org/gnu/grep/grep-3.11.tar.xz"
	grep_cmd_bin := exec.Command("curl", "-L", grep_bin_url, "-o", "binary.bin")
	err = grep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grep_src_url := "https://ftp.gnu.org/gnu/grep/grep-3.11.tar.xz"
	grep_cmd_src := exec.Command("curl", "-L", grep_src_url, "-o", "source.tar.gz")
	err = grep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grep_cmd_direct := exec.Command("./binary")
	err = grep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: wget")
	exec.Command("latte", "install", "wget").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
