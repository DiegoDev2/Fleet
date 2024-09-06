package main

// Zsh - UNIX shell (command interpreter)
// Homepage: https://www.zsh.org/

import (
	"fmt"
	
	"os/exec"
)

func installZsh() {
	// Método 1: Descargar y extraer .tar.gz
	zsh_tar_url := "https://downloads.sourceforge.net/project/zsh/zsh/5.9/zsh-5.9.tar.xz"
	zsh_cmd_tar := exec.Command("curl", "-L", zsh_tar_url, "-o", "package.tar.gz")
	err := zsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zsh_zip_url := "https://downloads.sourceforge.net/project/zsh/zsh/5.9/zsh-5.9.tar.xz"
	zsh_cmd_zip := exec.Command("curl", "-L", zsh_zip_url, "-o", "package.zip")
	err = zsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zsh_bin_url := "https://downloads.sourceforge.net/project/zsh/zsh/5.9/zsh-5.9.tar.xz"
	zsh_cmd_bin := exec.Command("curl", "-L", zsh_bin_url, "-o", "binary.bin")
	err = zsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zsh_src_url := "https://downloads.sourceforge.net/project/zsh/zsh/5.9/zsh-5.9.tar.xz"
	zsh_cmd_src := exec.Command("curl", "-L", zsh_src_url, "-o", "source.tar.gz")
	err = zsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zsh_cmd_direct := exec.Command("./binary")
	err = zsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}
