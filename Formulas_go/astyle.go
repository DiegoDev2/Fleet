package main

// Astyle - Source code beautifier for C, C++, C#, and Java
// Homepage: https://astyle.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAstyle() {
	// Método 1: Descargar y extraer .tar.gz
	astyle_tar_url := "https://downloads.sourceforge.net/project/astyle/astyle/astyle%203.6/astyle-3.6.tar.bz2"
	astyle_cmd_tar := exec.Command("curl", "-L", astyle_tar_url, "-o", "package.tar.gz")
	err := astyle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	astyle_zip_url := "https://downloads.sourceforge.net/project/astyle/astyle/astyle%203.6/astyle-3.6.tar.bz2"
	astyle_cmd_zip := exec.Command("curl", "-L", astyle_zip_url, "-o", "package.zip")
	err = astyle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	astyle_bin_url := "https://downloads.sourceforge.net/project/astyle/astyle/astyle%203.6/astyle-3.6.tar.bz2"
	astyle_cmd_bin := exec.Command("curl", "-L", astyle_bin_url, "-o", "binary.bin")
	err = astyle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	astyle_src_url := "https://downloads.sourceforge.net/project/astyle/astyle/astyle%203.6/astyle-3.6.tar.bz2"
	astyle_cmd_src := exec.Command("curl", "-L", astyle_src_url, "-o", "source.tar.gz")
	err = astyle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	astyle_cmd_direct := exec.Command("./binary")
	err = astyle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
