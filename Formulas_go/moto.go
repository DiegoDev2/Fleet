package main

// Moto - Mock AWS services
// Homepage: http://getmoto.org/

import (
	"fmt"
	
	"os/exec"
)

func installMoto() {
	// Método 1: Descargar y extraer .tar.gz
	moto_tar_url := "https://files.pythonhosted.org/packages/2e/0b/b5fce9229d252ea2a8cf4ddb08a847e99441d6feaa75f57121ab33c6d29f/moto-5.0.13.tar.gz"
	moto_cmd_tar := exec.Command("curl", "-L", moto_tar_url, "-o", "package.tar.gz")
	err := moto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moto_zip_url := "https://files.pythonhosted.org/packages/2e/0b/b5fce9229d252ea2a8cf4ddb08a847e99441d6feaa75f57121ab33c6d29f/moto-5.0.13.zip"
	moto_cmd_zip := exec.Command("curl", "-L", moto_zip_url, "-o", "package.zip")
	err = moto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moto_bin_url := "https://files.pythonhosted.org/packages/2e/0b/b5fce9229d252ea2a8cf4ddb08a847e99441d6feaa75f57121ab33c6d29f/moto-5.0.13.bin"
	moto_cmd_bin := exec.Command("curl", "-L", moto_bin_url, "-o", "binary.bin")
	err = moto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moto_src_url := "https://files.pythonhosted.org/packages/2e/0b/b5fce9229d252ea2a8cf4ddb08a847e99441d6feaa75f57121ab33c6d29f/moto-5.0.13.src.tar.gz"
	moto_cmd_src := exec.Command("curl", "-L", moto_src_url, "-o", "source.tar.gz")
	err = moto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moto_cmd_direct := exec.Command("./binary")
	err = moto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
