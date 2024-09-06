package main

// Pocsuite3 - Open-sourced remote vulnerability testing framework
// Homepage: https://pocsuite.org/

import (
	"fmt"
	
	"os/exec"
)

func installPocsuite3() {
	// Método 1: Descargar y extraer .tar.gz
	pocsuite3_tar_url := "https://files.pythonhosted.org/packages/0f/05/b17921332ab312c04ccc67b3d01a0d4318a4d45eb0315531f66d41a89639/pocsuite3-2.0.8.tar.gz"
	pocsuite3_cmd_tar := exec.Command("curl", "-L", pocsuite3_tar_url, "-o", "package.tar.gz")
	err := pocsuite3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pocsuite3_zip_url := "https://files.pythonhosted.org/packages/0f/05/b17921332ab312c04ccc67b3d01a0d4318a4d45eb0315531f66d41a89639/pocsuite3-2.0.8.zip"
	pocsuite3_cmd_zip := exec.Command("curl", "-L", pocsuite3_zip_url, "-o", "package.zip")
	err = pocsuite3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pocsuite3_bin_url := "https://files.pythonhosted.org/packages/0f/05/b17921332ab312c04ccc67b3d01a0d4318a4d45eb0315531f66d41a89639/pocsuite3-2.0.8.bin"
	pocsuite3_cmd_bin := exec.Command("curl", "-L", pocsuite3_bin_url, "-o", "binary.bin")
	err = pocsuite3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pocsuite3_src_url := "https://files.pythonhosted.org/packages/0f/05/b17921332ab312c04ccc67b3d01a0d4318a4d45eb0315531f66d41a89639/pocsuite3-2.0.8.src.tar.gz"
	pocsuite3_cmd_src := exec.Command("curl", "-L", pocsuite3_src_url, "-o", "source.tar.gz")
	err = pocsuite3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pocsuite3_cmd_direct := exec.Command("./binary")
	err = pocsuite3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
