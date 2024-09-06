package main

// Nu - Object-oriented, Lisp-like programming language
// Homepage: https://programming.nu/

import (
	"fmt"
	
	"os/exec"
)

func installNu() {
	// Método 1: Descargar y extraer .tar.gz
	nu_tar_url := "https://github.com/programming-nu/nu/archive/refs/tags/v2.3.0.tar.gz"
	nu_cmd_tar := exec.Command("curl", "-L", nu_tar_url, "-o", "package.tar.gz")
	err := nu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nu_zip_url := "https://github.com/programming-nu/nu/archive/refs/tags/v2.3.0.zip"
	nu_cmd_zip := exec.Command("curl", "-L", nu_zip_url, "-o", "package.zip")
	err = nu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nu_bin_url := "https://github.com/programming-nu/nu/archive/refs/tags/v2.3.0.bin"
	nu_cmd_bin := exec.Command("curl", "-L", nu_bin_url, "-o", "binary.bin")
	err = nu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nu_src_url := "https://github.com/programming-nu/nu/archive/refs/tags/v2.3.0.src.tar.gz"
	nu_cmd_src := exec.Command("curl", "-L", nu_src_url, "-o", "source.tar.gz")
	err = nu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nu_cmd_direct := exec.Command("./binary")
	err = nu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: gnustep-make")
	exec.Command("latte", "install", "gnustep-make").Run()
	fmt.Println("Instalando dependencia: gnustep-base")
	exec.Command("latte", "install", "gnustep-base").Run()
	fmt.Println("Instalando dependencia: libobjc2")
	exec.Command("latte", "install", "libobjc2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: libffi")
	exec.Command("latte", "install", "libffi").Run()
}
