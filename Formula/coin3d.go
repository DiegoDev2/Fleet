package main

// Coin3d - Open Inventor 2.1 API implementation (Coin) with Python bindings (Pivy)
// Homepage: https://coin3d.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installCoin3d() {
	// Método 1: Descargar y extraer .tar.gz
	coin3d_tar_url := "https://github.com/coin3d/coin/releases/download/v4.0.2/coin-4.0.2-src.zip"
	coin3d_cmd_tar := exec.Command("curl", "-L", coin3d_tar_url, "-o", "package.tar.gz")
	err := coin3d_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coin3d_zip_url := "https://github.com/coin3d/coin/releases/download/v4.0.2/coin-4.0.2-src.zip"
	coin3d_cmd_zip := exec.Command("curl", "-L", coin3d_zip_url, "-o", "package.zip")
	err = coin3d_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coin3d_bin_url := "https://github.com/coin3d/coin/releases/download/v4.0.2/coin-4.0.2-src.zip"
	coin3d_cmd_bin := exec.Command("curl", "-L", coin3d_bin_url, "-o", "binary.bin")
	err = coin3d_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coin3d_src_url := "https://github.com/coin3d/coin/releases/download/v4.0.2/coin-4.0.2-src.zip"
	coin3d_cmd_src := exec.Command("curl", "-L", coin3d_src_url, "-o", "source.tar.gz")
	err = coin3d_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coin3d_cmd_direct := exec.Command("./binary")
	err = coin3d_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: pyside@2")
	exec.Command("latte", "install", "pyside@2").Run()
	fmt.Println("Instalando dependencia: python@3.10")
	exec.Command("latte", "install", "python@3.10").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
