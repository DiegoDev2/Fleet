package main

// Lpeg - Parsing Expression Grammars For Lua
// Homepage: https://www.inf.puc-rio.br/~roberto/lpeg/

import (
	"fmt"
	
	"os/exec"
)

func installLpeg() {
	// Método 1: Descargar y extraer .tar.gz
	lpeg_tar_url := "https://www.inf.puc-rio.br/~roberto/lpeg/lpeg-1.1.0.tar.gz"
	lpeg_cmd_tar := exec.Command("curl", "-L", lpeg_tar_url, "-o", "package.tar.gz")
	err := lpeg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lpeg_zip_url := "https://www.inf.puc-rio.br/~roberto/lpeg/lpeg-1.1.0.zip"
	lpeg_cmd_zip := exec.Command("curl", "-L", lpeg_zip_url, "-o", "package.zip")
	err = lpeg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lpeg_bin_url := "https://www.inf.puc-rio.br/~roberto/lpeg/lpeg-1.1.0.bin"
	lpeg_cmd_bin := exec.Command("curl", "-L", lpeg_bin_url, "-o", "binary.bin")
	err = lpeg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lpeg_src_url := "https://www.inf.puc-rio.br/~roberto/lpeg/lpeg-1.1.0.src.tar.gz"
	lpeg_cmd_src := exec.Command("curl", "-L", lpeg_src_url, "-o", "source.tar.gz")
	err = lpeg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lpeg_cmd_direct := exec.Command("./binary")
	err = lpeg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
}
