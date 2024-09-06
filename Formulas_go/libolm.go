package main

// Libolm - Implementation of the Double Ratchet cryptographic ratchet
// Homepage: https://gitlab.matrix.org/matrix-org/olm

import (
	"fmt"
	
	"os/exec"
)

func installLibolm() {
	// Método 1: Descargar y extraer .tar.gz
	libolm_tar_url := "https://gitlab.matrix.org/matrix-org/olm/-/archive/3.2.16/olm-3.2.16.tar.gz"
	libolm_cmd_tar := exec.Command("curl", "-L", libolm_tar_url, "-o", "package.tar.gz")
	err := libolm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libolm_zip_url := "https://gitlab.matrix.org/matrix-org/olm/-/archive/3.2.16/olm-3.2.16.zip"
	libolm_cmd_zip := exec.Command("curl", "-L", libolm_zip_url, "-o", "package.zip")
	err = libolm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libolm_bin_url := "https://gitlab.matrix.org/matrix-org/olm/-/archive/3.2.16/olm-3.2.16.bin"
	libolm_cmd_bin := exec.Command("curl", "-L", libolm_bin_url, "-o", "binary.bin")
	err = libolm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libolm_src_url := "https://gitlab.matrix.org/matrix-org/olm/-/archive/3.2.16/olm-3.2.16.src.tar.gz"
	libolm_cmd_src := exec.Command("curl", "-L", libolm_src_url, "-o", "source.tar.gz")
	err = libolm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libolm_cmd_direct := exec.Command("./binary")
	err = libolm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
