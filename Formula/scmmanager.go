package main

// ScmManager - Manage Git, Mercurial, and Subversion repos over HTTP
// Homepage: https://www.scm-manager.org

import (
	"fmt"
	
	"os/exec"
)

func installScmManager() {
	// Método 1: Descargar y extraer .tar.gz
	scmmanager_tar_url := "https://packages.scm-manager.org/repository/releases/sonia/scm/packaging/unix/3.4.1/unix-3.4.1.tar.gz"
	scmmanager_cmd_tar := exec.Command("curl", "-L", scmmanager_tar_url, "-o", "package.tar.gz")
	err := scmmanager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scmmanager_zip_url := "https://packages.scm-manager.org/repository/releases/sonia/scm/packaging/unix/3.4.1/unix-3.4.1.zip"
	scmmanager_cmd_zip := exec.Command("curl", "-L", scmmanager_zip_url, "-o", "package.zip")
	err = scmmanager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scmmanager_bin_url := "https://packages.scm-manager.org/repository/releases/sonia/scm/packaging/unix/3.4.1/unix-3.4.1.bin"
	scmmanager_cmd_bin := exec.Command("curl", "-L", scmmanager_bin_url, "-o", "binary.bin")
	err = scmmanager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scmmanager_src_url := "https://packages.scm-manager.org/repository/releases/sonia/scm/packaging/unix/3.4.1/unix-3.4.1.src.tar.gz"
	scmmanager_cmd_src := exec.Command("curl", "-L", scmmanager_src_url, "-o", "source.tar.gz")
	err = scmmanager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scmmanager_cmd_direct := exec.Command("./binary")
	err = scmmanager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jsvc")
	exec.Command("latte", "install", "jsvc").Run()
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
