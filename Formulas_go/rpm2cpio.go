package main

// Rpm2cpio - Tool to convert RPM package to CPIO archive
// Homepage: https://svnweb.freebsd.org/ports/head/archivers/rpm2cpio/

import (
	"fmt"
	
	"os/exec"
)

func installRpm2cpio() {
	// Método 1: Descargar y extraer .tar.gz
	rpm2cpio_tar_url := "https://svnweb.freebsd.org/ports/head/archivers/rpm2cpio/files/rpm2cpio?revision=408590&view=co"
	rpm2cpio_cmd_tar := exec.Command("curl", "-L", rpm2cpio_tar_url, "-o", "package.tar.gz")
	err := rpm2cpio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rpm2cpio_zip_url := "https://svnweb.freebsd.org/ports/head/archivers/rpm2cpio/files/rpm2cpio?revision=408590&view=co"
	rpm2cpio_cmd_zip := exec.Command("curl", "-L", rpm2cpio_zip_url, "-o", "package.zip")
	err = rpm2cpio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rpm2cpio_bin_url := "https://svnweb.freebsd.org/ports/head/archivers/rpm2cpio/files/rpm2cpio?revision=408590&view=co"
	rpm2cpio_cmd_bin := exec.Command("curl", "-L", rpm2cpio_bin_url, "-o", "binary.bin")
	err = rpm2cpio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rpm2cpio_src_url := "https://svnweb.freebsd.org/ports/head/archivers/rpm2cpio/files/rpm2cpio?revision=408590&view=co"
	rpm2cpio_cmd_src := exec.Command("curl", "-L", rpm2cpio_src_url, "-o", "source.tar.gz")
	err = rpm2cpio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rpm2cpio_cmd_direct := exec.Command("./binary")
	err = rpm2cpio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
