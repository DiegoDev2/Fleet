package main

// SaneBackends - Backends for scanner access
// Homepage: http://www.sane-project.org/

import (
	"fmt"
	
	"os/exec"
)

func installSaneBackends() {
	// Método 1: Descargar y extraer .tar.gz
	sanebackends_tar_url := "https://gitlab.com/sane-project/backends/uploads/83bdbb6c9a115184c2d48f1fdc6847db/sane-backends-1.3.1.tar.gz"
	sanebackends_cmd_tar := exec.Command("curl", "-L", sanebackends_tar_url, "-o", "package.tar.gz")
	err := sanebackends_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sanebackends_zip_url := "https://gitlab.com/sane-project/backends/uploads/83bdbb6c9a115184c2d48f1fdc6847db/sane-backends-1.3.1.zip"
	sanebackends_cmd_zip := exec.Command("curl", "-L", sanebackends_zip_url, "-o", "package.zip")
	err = sanebackends_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sanebackends_bin_url := "https://gitlab.com/sane-project/backends/uploads/83bdbb6c9a115184c2d48f1fdc6847db/sane-backends-1.3.1.bin"
	sanebackends_cmd_bin := exec.Command("curl", "-L", sanebackends_bin_url, "-o", "binary.bin")
	err = sanebackends_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sanebackends_src_url := "https://gitlab.com/sane-project/backends/uploads/83bdbb6c9a115184c2d48f1fdc6847db/sane-backends-1.3.1.src.tar.gz"
	sanebackends_cmd_src := exec.Command("curl", "-L", sanebackends_src_url, "-o", "source.tar.gz")
	err = sanebackends_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sanebackends_cmd_direct := exec.Command("./binary")
	err = sanebackends_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: net-snmp")
exec.Command("latte", "install", "net-snmp")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
