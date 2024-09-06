package main

// Snapcraft - Package any app for every Linux desktop, server, cloud or device
// Homepage: https://snapcraft.io/

import (
	"fmt"
	
	"os/exec"
)

func installSnapcraft() {
	// Método 1: Descargar y extraer .tar.gz
	snapcraft_tar_url := "https://github.com/snapcore/snapcraft/archive/refs/tags/8.3.2.tar.gz"
	snapcraft_cmd_tar := exec.Command("curl", "-L", snapcraft_tar_url, "-o", "package.tar.gz")
	err := snapcraft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snapcraft_zip_url := "https://github.com/snapcore/snapcraft/archive/refs/tags/8.3.2.zip"
	snapcraft_cmd_zip := exec.Command("curl", "-L", snapcraft_zip_url, "-o", "package.zip")
	err = snapcraft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snapcraft_bin_url := "https://github.com/snapcore/snapcraft/archive/refs/tags/8.3.2.bin"
	snapcraft_cmd_bin := exec.Command("curl", "-L", snapcraft_bin_url, "-o", "binary.bin")
	err = snapcraft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snapcraft_src_url := "https://github.com/snapcore/snapcraft/archive/refs/tags/8.3.2.src.tar.gz"
	snapcraft_cmd_src := exec.Command("curl", "-L", snapcraft_src_url, "-o", "source.tar.gz")
	err = snapcraft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snapcraft_cmd_direct := exec.Command("./binary")
	err = snapcraft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgit2@1.7")
	exec.Command("latte", "install", "libgit2@1.7").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: lxc")
	exec.Command("latte", "install", "lxc").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: snap")
	exec.Command("latte", "install", "snap").Run()
	fmt.Println("Instalando dependencia: xdelta")
	exec.Command("latte", "install", "xdelta").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: apt")
	exec.Command("latte", "install", "apt").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
}
