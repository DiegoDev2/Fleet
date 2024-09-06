package main

// Bitlbee - IRC to other chat networks gateway
// Homepage: https://www.bitlbee.org/

import (
	"fmt"
	
	"os/exec"
)

func installBitlbee() {
	// Método 1: Descargar y extraer .tar.gz
	bitlbee_tar_url := "https://get.bitlbee.org/src/bitlbee-3.6.tar.gz"
	bitlbee_cmd_tar := exec.Command("curl", "-L", bitlbee_tar_url, "-o", "package.tar.gz")
	err := bitlbee_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bitlbee_zip_url := "https://get.bitlbee.org/src/bitlbee-3.6.zip"
	bitlbee_cmd_zip := exec.Command("curl", "-L", bitlbee_zip_url, "-o", "package.zip")
	err = bitlbee_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bitlbee_bin_url := "https://get.bitlbee.org/src/bitlbee-3.6.bin"
	bitlbee_cmd_bin := exec.Command("curl", "-L", bitlbee_bin_url, "-o", "binary.bin")
	err = bitlbee_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bitlbee_src_url := "https://get.bitlbee.org/src/bitlbee-3.6.src.tar.gz"
	bitlbee_cmd_src := exec.Command("curl", "-L", bitlbee_src_url, "-o", "source.tar.gz")
	err = bitlbee_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bitlbee_cmd_direct := exec.Command("./binary")
	err = bitlbee_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
