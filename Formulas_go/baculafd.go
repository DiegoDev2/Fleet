package main

// BaculaFd - Network backup solution
// Homepage: https://www.bacula.org/

import (
	"fmt"
	
	"os/exec"
)

func installBaculaFd() {
	// Método 1: Descargar y extraer .tar.gz
	baculafd_tar_url := "https://downloads.sourceforge.net/project/bacula/bacula/15.0.2/bacula-15.0.2.tar.gz"
	baculafd_cmd_tar := exec.Command("curl", "-L", baculafd_tar_url, "-o", "package.tar.gz")
	err := baculafd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	baculafd_zip_url := "https://downloads.sourceforge.net/project/bacula/bacula/15.0.2/bacula-15.0.2.zip"
	baculafd_cmd_zip := exec.Command("curl", "-L", baculafd_zip_url, "-o", "package.zip")
	err = baculafd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	baculafd_bin_url := "https://downloads.sourceforge.net/project/bacula/bacula/15.0.2/bacula-15.0.2.bin"
	baculafd_cmd_bin := exec.Command("curl", "-L", baculafd_bin_url, "-o", "binary.bin")
	err = baculafd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	baculafd_src_url := "https://downloads.sourceforge.net/project/bacula/bacula/15.0.2/bacula-15.0.2.src.tar.gz"
	baculafd_cmd_src := exec.Command("curl", "-L", baculafd_src_url, "-o", "source.tar.gz")
	err = baculafd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	baculafd_cmd_direct := exec.Command("./binary")
	err = baculafd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
