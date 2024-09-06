package main

// UniversalCtags - Maintained ctags implementation
// Homepage: https://github.com/universal-ctags/ctags

import (
	"fmt"
	
	"os/exec"
)

func installUniversalCtags() {
	// Método 1: Descargar y extraer .tar.gz
	universalctags_tar_url := "https://github.com/universal-ctags/ctags/archive/refs/tags/p6.1.20240901.0.tar.gz"
	universalctags_cmd_tar := exec.Command("curl", "-L", universalctags_tar_url, "-o", "package.tar.gz")
	err := universalctags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	universalctags_zip_url := "https://github.com/universal-ctags/ctags/archive/refs/tags/p6.1.20240901.0.zip"
	universalctags_cmd_zip := exec.Command("curl", "-L", universalctags_zip_url, "-o", "package.zip")
	err = universalctags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	universalctags_bin_url := "https://github.com/universal-ctags/ctags/archive/refs/tags/p6.1.20240901.0.bin"
	universalctags_cmd_bin := exec.Command("curl", "-L", universalctags_bin_url, "-o", "binary.bin")
	err = universalctags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	universalctags_src_url := "https://github.com/universal-ctags/ctags/archive/refs/tags/p6.1.20240901.0.src.tar.gz"
	universalctags_cmd_src := exec.Command("curl", "-L", universalctags_src_url, "-o", "source.tar.gz")
	err = universalctags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	universalctags_cmd_direct := exec.Command("./binary")
	err = universalctags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: docutils")
exec.Command("latte", "install", "docutils")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
