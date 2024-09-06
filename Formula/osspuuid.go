package main

// OsspUuid - ISO-C API and CLI for generating UUIDs
// Homepage: http://www.ossp.org/pkg/lib/uuid/

import (
	"fmt"
	
	"os/exec"
)

func installOsspUuid() {
	// Método 1: Descargar y extraer .tar.gz
	osspuuid_tar_url := "https://deb.debian.org/debian/pool/main/o/ossp-uuid/ossp-uuid_1.6.2.orig.tar.gz"
	osspuuid_cmd_tar := exec.Command("curl", "-L", osspuuid_tar_url, "-o", "package.tar.gz")
	err := osspuuid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osspuuid_zip_url := "https://deb.debian.org/debian/pool/main/o/ossp-uuid/ossp-uuid_1.6.2.orig.zip"
	osspuuid_cmd_zip := exec.Command("curl", "-L", osspuuid_zip_url, "-o", "package.zip")
	err = osspuuid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osspuuid_bin_url := "https://deb.debian.org/debian/pool/main/o/ossp-uuid/ossp-uuid_1.6.2.orig.bin"
	osspuuid_cmd_bin := exec.Command("curl", "-L", osspuuid_bin_url, "-o", "binary.bin")
	err = osspuuid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osspuuid_src_url := "https://deb.debian.org/debian/pool/main/o/ossp-uuid/ossp-uuid_1.6.2.orig.src.tar.gz"
	osspuuid_cmd_src := exec.Command("curl", "-L", osspuuid_src_url, "-o", "source.tar.gz")
	err = osspuuid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osspuuid_cmd_direct := exec.Command("./binary")
	err = osspuuid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
