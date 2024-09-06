package main

// Credstash - Little utility for managing credentials in the cloud
// Homepage: https://github.com/fugue/credstash

import (
	"fmt"
	
	"os/exec"
)

func installCredstash() {
	// Método 1: Descargar y extraer .tar.gz
	credstash_tar_url := "https://files.pythonhosted.org/packages/b4/89/f929fda5fec87046873be2420a4c0cb40a82ab5e30c6d9cb22ddec41450b/credstash-1.17.1.tar.gz"
	credstash_cmd_tar := exec.Command("curl", "-L", credstash_tar_url, "-o", "package.tar.gz")
	err := credstash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	credstash_zip_url := "https://files.pythonhosted.org/packages/b4/89/f929fda5fec87046873be2420a4c0cb40a82ab5e30c6d9cb22ddec41450b/credstash-1.17.1.zip"
	credstash_cmd_zip := exec.Command("curl", "-L", credstash_zip_url, "-o", "package.zip")
	err = credstash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	credstash_bin_url := "https://files.pythonhosted.org/packages/b4/89/f929fda5fec87046873be2420a4c0cb40a82ab5e30c6d9cb22ddec41450b/credstash-1.17.1.bin"
	credstash_cmd_bin := exec.Command("curl", "-L", credstash_bin_url, "-o", "binary.bin")
	err = credstash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	credstash_src_url := "https://files.pythonhosted.org/packages/b4/89/f929fda5fec87046873be2420a4c0cb40a82ab5e30c6d9cb22ddec41450b/credstash-1.17.1.src.tar.gz"
	credstash_cmd_src := exec.Command("curl", "-L", credstash_src_url, "-o", "source.tar.gz")
	err = credstash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	credstash_cmd_direct := exec.Command("./binary")
	err = credstash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
