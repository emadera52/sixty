			<div class="col-md-10 comment">
				<br>
				<form action="/comments" method="POST">
					<input type="radio" name="category" value="1" {{.IsChk1}}> Adventurer&nbsp;
					<input type="radio" name="category" value="2" {{.IsChk2}}> Advocate&nbsp; 
					<input type="radio" name="category" value="3" {{.IsChk3}}> Provider&nbsp; 
					<input type="radio" name="category" value="4" {{.IsChk4}}> Supporter&nbsp; 
					<input type="radio" name="category" value="0" {{.IsChk0}}> Just Looking
					<div class="img-info">
					  <img src="static/img/base/info.png" alt="info" />
					  <p class="img-text"><em>Indicate or Update</em> your primary interest by selecting one of these options.</p>
					</div>
					<br>
					<textarea rows="6" name="cmnttext" placeholder="Enter your commments here."></textarea>
					{{ .xsrftoken }}
					<input type="hidden" name="cmntcat" value={{.CmntCat}}>
					<br>
					<input class="button" type="submit"
					value="Submit"/>
				</form>
			</div>
