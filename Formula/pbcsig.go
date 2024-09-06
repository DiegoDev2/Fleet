package main

// PbcSig - Signatures library
// Homepage: https://crypto.stanford.edu/pbc/sig/

import (
	"fmt"
	
	"os/exec"
)

func installPbcSig() {
	// Método 1: Descargar y extraer .tar.gz
	pbcsig_tar_url := "https://crypto.stanford.edu/pbc/sig/files/pbc_sig-0.0.8.tar.gz"
	pbcsig_cmd_tar := exec.Command("curl", "-L", pbcsig_tar_url, "-o", "package.tar.gz")
	err := pbcsig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pbcsig_zip_url := "https://crypto.stanford.edu/pbc/sig/files/pbc_sig-0.0.8.zip"
	pbcsig_cmd_zip := exec.Command("curl", "-L", pbcsig_zip_url, "-o", "package.zip")
	err = pbcsig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pbcsig_bin_url := "https://crypto.stanford.edu/pbc/sig/files/pbc_sig-0.0.8.bin"
	pbcsig_cmd_bin := exec.Command("curl", "-L", pbcsig_bin_url, "-o", "binary.bin")
	err = pbcsig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pbcsig_src_url := "https://crypto.stanford.edu/pbc/sig/files/pbc_sig-0.0.8.src.tar.gz"
	pbcsig_cmd_src := exec.Command("curl", "-L", pbcsig_src_url, "-o", "source.tar.gz")
	err = pbcsig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pbcsig_cmd_direct := exec.Command("./binary")
	err = pbcsig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: pbc")
	exec.Command("latte", "install", "pbc").Run()
}
