package main

// West - Zephyr meta-tool
// Homepage: https://github.com/zephyrproject-rtos/west

import (
	"fmt"
	
	"os/exec"
)

func installWest() {
	// Método 1: Descargar y extraer .tar.gz
	west_tar_url := "https://files.pythonhosted.org/packages/ee/7a/4c69c6a1054b319421d5acf028564bb1303ea9da42032a2000021d6495ee/west-1.2.0.tar.gz"
	west_cmd_tar := exec.Command("curl", "-L", west_tar_url, "-o", "package.tar.gz")
	err := west_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	west_zip_url := "https://files.pythonhosted.org/packages/ee/7a/4c69c6a1054b319421d5acf028564bb1303ea9da42032a2000021d6495ee/west-1.2.0.zip"
	west_cmd_zip := exec.Command("curl", "-L", west_zip_url, "-o", "package.zip")
	err = west_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	west_bin_url := "https://files.pythonhosted.org/packages/ee/7a/4c69c6a1054b319421d5acf028564bb1303ea9da42032a2000021d6495ee/west-1.2.0.bin"
	west_cmd_bin := exec.Command("curl", "-L", west_bin_url, "-o", "binary.bin")
	err = west_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	west_src_url := "https://files.pythonhosted.org/packages/ee/7a/4c69c6a1054b319421d5acf028564bb1303ea9da42032a2000021d6495ee/west-1.2.0.src.tar.gz"
	west_cmd_src := exec.Command("curl", "-L", west_src_url, "-o", "source.tar.gz")
	err = west_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	west_cmd_direct := exec.Command("./binary")
	err = west_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
