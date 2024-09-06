package main

// Pbc - Pairing-based cryptography
// Homepage: https://crypto.stanford.edu/pbc/

import (
	"fmt"
	
	"os/exec"
)

func installPbc() {
	// Método 1: Descargar y extraer .tar.gz
	pbc_tar_url := "https://crypto.stanford.edu/pbc/files/pbc-0.5.14.tar.gz"
	pbc_cmd_tar := exec.Command("curl", "-L", pbc_tar_url, "-o", "package.tar.gz")
	err := pbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pbc_zip_url := "https://crypto.stanford.edu/pbc/files/pbc-0.5.14.zip"
	pbc_cmd_zip := exec.Command("curl", "-L", pbc_zip_url, "-o", "package.zip")
	err = pbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pbc_bin_url := "https://crypto.stanford.edu/pbc/files/pbc-0.5.14.bin"
	pbc_cmd_bin := exec.Command("curl", "-L", pbc_bin_url, "-o", "binary.bin")
	err = pbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pbc_src_url := "https://crypto.stanford.edu/pbc/files/pbc-0.5.14.src.tar.gz"
	pbc_cmd_src := exec.Command("curl", "-L", pbc_src_url, "-o", "source.tar.gz")
	err = pbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pbc_cmd_direct := exec.Command("./binary")
	err = pbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}
