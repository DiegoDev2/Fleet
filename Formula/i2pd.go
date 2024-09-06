package main

// I2pd - Full-featured C++ implementation of I2P client
// Homepage: https://i2pd.website/

import (
	"fmt"
	
	"os/exec"
)

func installI2pd() {
	// Método 1: Descargar y extraer .tar.gz
	i2pd_tar_url := "https://github.com/PurpleI2P/i2pd/archive/refs/tags/2.53.1.tar.gz"
	i2pd_cmd_tar := exec.Command("curl", "-L", i2pd_tar_url, "-o", "package.tar.gz")
	err := i2pd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	i2pd_zip_url := "https://github.com/PurpleI2P/i2pd/archive/refs/tags/2.53.1.zip"
	i2pd_cmd_zip := exec.Command("curl", "-L", i2pd_zip_url, "-o", "package.zip")
	err = i2pd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	i2pd_bin_url := "https://github.com/PurpleI2P/i2pd/archive/refs/tags/2.53.1.bin"
	i2pd_cmd_bin := exec.Command("curl", "-L", i2pd_bin_url, "-o", "binary.bin")
	err = i2pd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	i2pd_src_url := "https://github.com/PurpleI2P/i2pd/archive/refs/tags/2.53.1.src.tar.gz"
	i2pd_cmd_src := exec.Command("curl", "-L", i2pd_src_url, "-o", "source.tar.gz")
	err = i2pd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	i2pd_cmd_direct := exec.Command("./binary")
	err = i2pd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: miniupnpc")
	exec.Command("latte", "install", "miniupnpc").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
