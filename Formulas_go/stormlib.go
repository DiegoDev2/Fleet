package main

// Stormlib - Library for handling Blizzard MPQ archives
// Homepage: http://www.zezula.net/en/mpq/stormlib.html

import (
	"fmt"
	
	"os/exec"
)

func installStormlib() {
	// Método 1: Descargar y extraer .tar.gz
	stormlib_tar_url := "https://github.com/ladislav-zezula/StormLib/archive/refs/tags/v9.26.tar.gz"
	stormlib_cmd_tar := exec.Command("curl", "-L", stormlib_tar_url, "-o", "package.tar.gz")
	err := stormlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stormlib_zip_url := "https://github.com/ladislav-zezula/StormLib/archive/refs/tags/v9.26.zip"
	stormlib_cmd_zip := exec.Command("curl", "-L", stormlib_zip_url, "-o", "package.zip")
	err = stormlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stormlib_bin_url := "https://github.com/ladislav-zezula/StormLib/archive/refs/tags/v9.26.bin"
	stormlib_cmd_bin := exec.Command("curl", "-L", stormlib_bin_url, "-o", "binary.bin")
	err = stormlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stormlib_src_url := "https://github.com/ladislav-zezula/StormLib/archive/refs/tags/v9.26.src.tar.gz"
	stormlib_cmd_src := exec.Command("curl", "-L", stormlib_src_url, "-o", "source.tar.gz")
	err = stormlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stormlib_cmd_direct := exec.Command("./binary")
	err = stormlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
