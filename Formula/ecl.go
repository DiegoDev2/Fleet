package main

// Ecl - Embeddable Common Lisp
// Homepage: https://ecl.common-lisp.dev

import (
	"fmt"
	
	"os/exec"
)

func installEcl() {
	// Método 1: Descargar y extraer .tar.gz
	ecl_tar_url := "https://ecl.common-lisp.dev/static/files/release/ecl-24.5.10.tgz"
	ecl_cmd_tar := exec.Command("curl", "-L", ecl_tar_url, "-o", "package.tar.gz")
	err := ecl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ecl_zip_url := "https://ecl.common-lisp.dev/static/files/release/ecl-24.5.10.tgz"
	ecl_cmd_zip := exec.Command("curl", "-L", ecl_zip_url, "-o", "package.zip")
	err = ecl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ecl_bin_url := "https://ecl.common-lisp.dev/static/files/release/ecl-24.5.10.tgz"
	ecl_cmd_bin := exec.Command("curl", "-L", ecl_bin_url, "-o", "binary.bin")
	err = ecl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ecl_src_url := "https://ecl.common-lisp.dev/static/files/release/ecl-24.5.10.tgz"
	ecl_cmd_src := exec.Command("curl", "-L", ecl_src_url, "-o", "source.tar.gz")
	err = ecl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ecl_cmd_direct := exec.Command("./binary")
	err = ecl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}
