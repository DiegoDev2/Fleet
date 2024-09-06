package main

// I2cTools - Heterogeneous set of I2C tools for Linux
// Homepage: https://i2c.wiki.kernel.org/index.php/I2C_Tools

import (
	"fmt"
	
	"os/exec"
)

func installI2cTools() {
	// Método 1: Descargar y extraer .tar.gz
	i2ctools_tar_url := "https://mirrors.edge.kernel.org/pub/software/utils/i2c-tools/i2c-tools-4.3.tar.xz"
	i2ctools_cmd_tar := exec.Command("curl", "-L", i2ctools_tar_url, "-o", "package.tar.gz")
	err := i2ctools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	i2ctools_zip_url := "https://mirrors.edge.kernel.org/pub/software/utils/i2c-tools/i2c-tools-4.3.tar.xz"
	i2ctools_cmd_zip := exec.Command("curl", "-L", i2ctools_zip_url, "-o", "package.zip")
	err = i2ctools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	i2ctools_bin_url := "https://mirrors.edge.kernel.org/pub/software/utils/i2c-tools/i2c-tools-4.3.tar.xz"
	i2ctools_cmd_bin := exec.Command("curl", "-L", i2ctools_bin_url, "-o", "binary.bin")
	err = i2ctools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	i2ctools_src_url := "https://mirrors.edge.kernel.org/pub/software/utils/i2c-tools/i2c-tools-4.3.tar.xz"
	i2ctools_cmd_src := exec.Command("curl", "-L", i2ctools_src_url, "-o", "source.tar.gz")
	err = i2ctools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	i2ctools_cmd_direct := exec.Command("./binary")
	err = i2ctools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
