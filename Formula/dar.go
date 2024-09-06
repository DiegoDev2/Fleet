package main

// Dar - Backup directory tree and files
// Homepage: http://dar.linux.free.fr/doc/index.html

import (
	"fmt"
	
	"os/exec"
)

func installDar() {
	// Método 1: Descargar y extraer .tar.gz
	dar_tar_url := "https://downloads.sourceforge.net/project/dar/dar/2.7.15/dar-2.7.15.tar.gz"
	dar_cmd_tar := exec.Command("curl", "-L", dar_tar_url, "-o", "package.tar.gz")
	err := dar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dar_zip_url := "https://downloads.sourceforge.net/project/dar/dar/2.7.15/dar-2.7.15.zip"
	dar_cmd_zip := exec.Command("curl", "-L", dar_zip_url, "-o", "package.zip")
	err = dar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dar_bin_url := "https://downloads.sourceforge.net/project/dar/dar/2.7.15/dar-2.7.15.bin"
	dar_cmd_bin := exec.Command("curl", "-L", dar_bin_url, "-o", "binary.bin")
	err = dar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dar_src_url := "https://downloads.sourceforge.net/project/dar/dar/2.7.15/dar-2.7.15.src.tar.gz"
	dar_cmd_src := exec.Command("curl", "-L", dar_src_url, "-o", "source.tar.gz")
	err = dar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dar_cmd_direct := exec.Command("./binary")
	err = dar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: argon2")
	exec.Command("latte", "install", "argon2").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
}
