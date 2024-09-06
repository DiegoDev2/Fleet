package main

// TaLib - Tools for market analysis
// Homepage: https://ta-lib.org/

import (
	"fmt"
	
	"os/exec"
)

func installTaLib() {
	// Método 1: Descargar y extraer .tar.gz
	talib_tar_url := "https://downloads.sourceforge.net/project/ta-lib/ta-lib/0.4.0/ta-lib-0.4.0-src.tar.gz"
	talib_cmd_tar := exec.Command("curl", "-L", talib_tar_url, "-o", "package.tar.gz")
	err := talib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	talib_zip_url := "https://downloads.sourceforge.net/project/ta-lib/ta-lib/0.4.0/ta-lib-0.4.0-src.zip"
	talib_cmd_zip := exec.Command("curl", "-L", talib_zip_url, "-o", "package.zip")
	err = talib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	talib_bin_url := "https://downloads.sourceforge.net/project/ta-lib/ta-lib/0.4.0/ta-lib-0.4.0-src.bin"
	talib_cmd_bin := exec.Command("curl", "-L", talib_bin_url, "-o", "binary.bin")
	err = talib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	talib_src_url := "https://downloads.sourceforge.net/project/ta-lib/ta-lib/0.4.0/ta-lib-0.4.0-src.src.tar.gz"
	talib_cmd_src := exec.Command("curl", "-L", talib_src_url, "-o", "source.tar.gz")
	err = talib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	talib_cmd_direct := exec.Command("./binary")
	err = talib_cmd_direct.Run()
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
}
