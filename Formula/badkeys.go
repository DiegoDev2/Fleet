package main

// Badkeys - Tool to find common vulnerabilities in cryptographic public keys
// Homepage: https://badkeys.info

import (
	"fmt"
	
	"os/exec"
)

func installBadkeys() {
	// Método 1: Descargar y extraer .tar.gz
	badkeys_tar_url := "https://files.pythonhosted.org/packages/88/fd/0b40be2d9d46befa087cc5ca494ebadf5777cb05a5ef6ee27577e82ae409/badkeys-0.0.11.tar.gz"
	badkeys_cmd_tar := exec.Command("curl", "-L", badkeys_tar_url, "-o", "package.tar.gz")
	err := badkeys_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	badkeys_zip_url := "https://files.pythonhosted.org/packages/88/fd/0b40be2d9d46befa087cc5ca494ebadf5777cb05a5ef6ee27577e82ae409/badkeys-0.0.11.zip"
	badkeys_cmd_zip := exec.Command("curl", "-L", badkeys_zip_url, "-o", "package.zip")
	err = badkeys_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	badkeys_bin_url := "https://files.pythonhosted.org/packages/88/fd/0b40be2d9d46befa087cc5ca494ebadf5777cb05a5ef6ee27577e82ae409/badkeys-0.0.11.bin"
	badkeys_cmd_bin := exec.Command("curl", "-L", badkeys_bin_url, "-o", "binary.bin")
	err = badkeys_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	badkeys_src_url := "https://files.pythonhosted.org/packages/88/fd/0b40be2d9d46befa087cc5ca494ebadf5777cb05a5ef6ee27577e82ae409/badkeys-0.0.11.src.tar.gz"
	badkeys_cmd_src := exec.Command("curl", "-L", badkeys_src_url, "-o", "source.tar.gz")
	err = badkeys_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	badkeys_cmd_direct := exec.Command("./binary")
	err = badkeys_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libmpc")
	exec.Command("latte", "install", "libmpc").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
