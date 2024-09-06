package main

// Krb5 - Network authentication protocol
// Homepage: https://web.mit.edu/kerberos/

import (
	"fmt"
	
	"os/exec"
)

func installKrb5() {
	// Método 1: Descargar y extraer .tar.gz
	krb5_tar_url := "https://kerberos.org/dist/krb5/1.21/krb5-1.21.3.tar.gz"
	krb5_cmd_tar := exec.Command("curl", "-L", krb5_tar_url, "-o", "package.tar.gz")
	err := krb5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	krb5_zip_url := "https://kerberos.org/dist/krb5/1.21/krb5-1.21.3.zip"
	krb5_cmd_zip := exec.Command("curl", "-L", krb5_zip_url, "-o", "package.zip")
	err = krb5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	krb5_bin_url := "https://kerberos.org/dist/krb5/1.21/krb5-1.21.3.bin"
	krb5_cmd_bin := exec.Command("curl", "-L", krb5_bin_url, "-o", "binary.bin")
	err = krb5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	krb5_src_url := "https://kerberos.org/dist/krb5/1.21/krb5-1.21.3.src.tar.gz"
	krb5_cmd_src := exec.Command("curl", "-L", krb5_src_url, "-o", "source.tar.gz")
	err = krb5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	krb5_cmd_direct := exec.Command("./binary")
	err = krb5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
