package main

// I2p - Anonymous overlay network - a network within a network
// Homepage: https://geti2p.net

import (
	"fmt"
	
	"os/exec"
)

func installI2p() {
	// Método 1: Descargar y extraer .tar.gz
	i2p_tar_url := "https://github.com/i2p/i2p.i2p/releases/download/i2p-2.6.1/i2psource_2.6.1.tar.bz2"
	i2p_cmd_tar := exec.Command("curl", "-L", i2p_tar_url, "-o", "package.tar.gz")
	err := i2p_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	i2p_zip_url := "https://github.com/i2p/i2p.i2p/releases/download/i2p-2.6.1/i2psource_2.6.1.tar.bz2"
	i2p_cmd_zip := exec.Command("curl", "-L", i2p_zip_url, "-o", "package.zip")
	err = i2p_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	i2p_bin_url := "https://github.com/i2p/i2p.i2p/releases/download/i2p-2.6.1/i2psource_2.6.1.tar.bz2"
	i2p_cmd_bin := exec.Command("curl", "-L", i2p_bin_url, "-o", "binary.bin")
	err = i2p_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	i2p_src_url := "https://github.com/i2p/i2p.i2p/releases/download/i2p-2.6.1/i2psource_2.6.1.tar.bz2"
	i2p_cmd_src := exec.Command("curl", "-L", i2p_src_url, "-o", "source.tar.gz")
	err = i2p_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	i2p_cmd_direct := exec.Command("./binary")
	err = i2p_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: java-service-wrapper")
exec.Command("latte", "install", "java-service-wrapper")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
