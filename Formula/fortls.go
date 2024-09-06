package main

// Fortls - Fortran language server
// Homepage: https://fortls.fortran-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installFortls() {
	// Método 1: Descargar y extraer .tar.gz
	fortls_tar_url := "https://files.pythonhosted.org/packages/f1/82/b0f91372538de824bccb5e4fe8936e47f6771dbd700a74d35e19045050b5/fortls-3.1.2.tar.gz"
	fortls_cmd_tar := exec.Command("curl", "-L", fortls_tar_url, "-o", "package.tar.gz")
	err := fortls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fortls_zip_url := "https://files.pythonhosted.org/packages/f1/82/b0f91372538de824bccb5e4fe8936e47f6771dbd700a74d35e19045050b5/fortls-3.1.2.zip"
	fortls_cmd_zip := exec.Command("curl", "-L", fortls_zip_url, "-o", "package.zip")
	err = fortls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fortls_bin_url := "https://files.pythonhosted.org/packages/f1/82/b0f91372538de824bccb5e4fe8936e47f6771dbd700a74d35e19045050b5/fortls-3.1.2.bin"
	fortls_cmd_bin := exec.Command("curl", "-L", fortls_bin_url, "-o", "binary.bin")
	err = fortls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fortls_src_url := "https://files.pythonhosted.org/packages/f1/82/b0f91372538de824bccb5e4fe8936e47f6771dbd700a74d35e19045050b5/fortls-3.1.2.src.tar.gz"
	fortls_cmd_src := exec.Command("curl", "-L", fortls_src_url, "-o", "source.tar.gz")
	err = fortls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fortls_cmd_direct := exec.Command("./binary")
	err = fortls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
